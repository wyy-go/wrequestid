package wrequestid

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sony/sonyflake"
)


const headerXRequestID = "X-Request-ID"

type ctxRequestIDKey struct{}
type Option func(*Options)

func WithRequestIDHeader(header string) Option {
	return func(o *Options) {
		o.headerRequestID = header
	}
}

func WithGenerator(gen func() string) Option {
	return func(o *Options) {
		o.generator = gen
	}
}

// Options defines the config for RequestID middleware
type Options struct {
	generator func() string
	headerRequestID string
}

// New initializes the RequestID middleware.
func New(opts ...Option) gin.HandlerFunc {
	options := Options {
		headerRequestID: headerXRequestID,
		generator: defaultGenerator,
	}
	for _, opt := range opts {
		opt(&options)
	}

	return func(c *gin.Context) {
		// Get id from request
		rid := c.GetHeader(options.headerRequestID)
		if rid == "" {
			rid = options.generator()
		}

		// Set the id to ensure that the requestid is in the response
		c.Header(options.headerRequestID, rid)

		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, ctxRequestIDKey{}, rid)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}


func FromRequestID(ctx context.Context) string {
	rid, ok := ctx.Value(ctxRequestIDKey{}).(string)
	if !ok {
		return ""
	}
	return rid
}

// GetRequestID returns the request identifier
func GetRequestID(c *gin.Context) string {
	return FromRequestID(c.Request.Context())
}

var sf *sonyflake.Sonyflake
func init() {
	var st sonyflake.Settings
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
}

func defaultGenerator() string {
	id,_ := sf.NextID()
	return fmt.Sprintf("%d",id)
}