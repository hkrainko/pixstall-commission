package middleware

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
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

func (j *JWTPayloadsExtractor) ExtractPayloadsFromJWT(c *gin.Context) {
	fmt.Println("in ExtractIDFromJWT")
	jwtToken := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	print(jwtToken)

	ss := strings.Split(jwtToken, ".")

	if len(ss) != 3 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	b, err := base64.RawStdEncoding.DecodeString(ss[1])
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