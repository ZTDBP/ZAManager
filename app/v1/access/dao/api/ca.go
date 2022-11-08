package api

import (
	"crypto/x509/pkix"
	"encoding/hex"
	"os"
	"sync"
	"time"

	"github.com/ZTDBP/ZAManager/pkg/confer"
	"github.com/ZTDBP/ZAManager/pkg/logger"
	"github.com/ZTDBP/cfssl/helpers"
	"github.com/ZTDBP/zaca-sdk/caclient"
	"github.com/ZTDBP/zaca-sdk/keygen"
	"github.com/ZTDBP/zaca-sdk/pkg/attrmgr"
	"github.com/ZTDBP/zaca-sdk/pkg/spiffe"
	"github.com/gin-gonic/gin"
)

var once sync.Once
var caClient *caclient.CAInstance

type SentinelSign struct {
	CaPEM     string
	CertPEM   string
	KeyPEM    string
	Sn        string
	Aki       string
	ExpiredAt time.Time
}

func getCaClient() *caclient.CAInstance {
	once.Do(func() {
		if caClient == nil {
			cfg := confer.GlobalConfig().CA
			caClient = caclient.NewCAI(
				caclient.WithCAServer(caclient.RoleDefault /*哨兵*/, cfg.SignURL),
				caclient.WithAuthKey(cfg.AuthKey),
			)
		}
	})
	return caClient
}

func ApplySign(c *gin.Context, attrs map[string]interface{}, uniqueID, cn, host string, expiredAt time.Time) (sentinelSign SentinelSign, err error) {
	client := getCaClient()
	mgr, err := client.NewCertManager()
	if err != nil {
		return
	}
	// CA PEM
	caPEMBytes, err := mgr.CACertsPEM()
	if err != nil {
		logger.Errorf(c, "mgr.CACertsPEM() err : %v", err)
		return
	}
	caPEM := string(caPEMBytes)
	// KEY PEM
	_, keyPEMBytes, _ := keygen.GenKey(keygen.EcdsaSigAlg)
	// 证书扩展字段
	attr := attrmgr.New()
	ext, _ := attr.ToPkixExtension(&attrmgr.Attributes{
		// 注入参数 Map[string]interface{}
		Attrs: attrs,
	})
	// gen csr
	csrPEM, _ := keygen.GenCustomExtendCSR(keyPEMBytes, &spiffe.IDGIdentity{
		SiteID:    os.Getenv("IDG_SITEUID"), /* Site 标识 */
		ClusterID: os.Getenv("IDG_CLUSTERUID"),
		UniqueID:  uniqueID,
	}, &keygen.CertOptions{ /* 通常为固定值 */
		CN:   cn,
		Host: host,
	}, []pkix.Extension{ext} /* 注入扩展字段 */)
	// get cert
	certPEMBytes, err := mgr.SignPEM(csrPEM, uniqueID)
	if err != nil {
		logger.Errorf(c, "mgr.SignPEM() err : %v", err)
		return
	}
	certPEM := string(certPEMBytes)
	cert, err := helpers.ParseCertificatePEM(certPEMBytes)
	if err != nil {
		logger.Errorf(c, "helpers.ParseCertificatePEM() err : %v", err)
		return
	}
	sentinelSign = SentinelSign{
		CaPEM:     caPEM,
		CertPEM:   certPEM,
		KeyPEM:    string(keyPEMBytes),
		Sn:        cert.SerialNumber.String(),
		Aki:       hex.EncodeToString(cert.AuthorityKeyId),
		ExpiredAt: expiredAt,
		//ExpiredAt: cert.NotAfter,
	}
	return
}
