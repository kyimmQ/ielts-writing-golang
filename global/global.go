package global

import (
	"github.com/kyimmQ/ielts-writing-golang/pkg/settings"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	Config  settings.Config
	MongoDB *mongo.Client
)
