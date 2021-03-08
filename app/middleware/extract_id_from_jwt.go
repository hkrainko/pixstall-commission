package middleware

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type JWTPayloadsExtractor struct {
	Keys []string
}

func NewJWTPayloadsExtractor(keys []string) *JWTPayloadsExtractor {
	return &JWTPayloadsExtractor{
		Keys: keys,
	}
}

func (j *JWTPayloadsExtractor) ExtractPayloadsFromJWTInHeader(c *gin.Context) {
	jwtToken := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	payload := j.getPayloadPartFromFullToken(jwtToken)
	if payload == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	j.passTokenToNext(c, *payload)
}

func (j *JWTPayloadsExtractor) ExtractPayloadsFromJWTInQuery(c *gin.Context) {
	jwtToken := c.Query("access_token")
	payload := j.getPayloadPartFromFullToken(jwtToken)
	if payload == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	j.passTokenToNext(c, *payload)
}


// Private
func (j *JWTPayloadsExtractor) getPayloadPartFromFullToken(token string) *string {
	ss := strings.Split(token, ".")
	if len(ss) != 3 {
		return nil
	}
	return &ss[1]
}

func (j *JWTPayloadsExtractor) passTokenToNext(c *gin.Context, token string) {
	b, err := base64.RawStdEncoding.DecodeString(token)
	if b == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var nMap map[string]interface{}
	err = json.Unmarshal(b, &nMap)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	for _, key := range j.Keys {
		if nMap[key] != nil {
			if v, ok := nMap[key].(string); ok {
				c.Set(key, v)
			}
		}
	}
	c.Next()
}

type JWTPayload map[string]string