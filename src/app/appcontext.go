package app

import (
	"context"

	"github.com/go-playground/validator/v10"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type appContext struct {
	config    *Configuration
	gorm      *gorm.DB
	validator *validator.Validate
	tracer    trace.Tracer
}

var appCtx appContext

func Init(ctx context.Context) error {
	// Init Config
	config, err := InitConfig(ctx)
	if err != nil {
		return err
	}

	// TODO: Add Gorm
	// Gorm Config
	// gormConfig := gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	// dbConfig := config.MySQL
	// gorm, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbConfig.Username, dbConfig.Password, dbConfig.ConnURI, dbConfig.Database)), &gormConfig)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// Set Gorm Tracing
	// if err := gorm.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
	// 	log.Panic(err)
	// }

	appCtx = appContext{
		// gorm:      gorm,
		config:    config,
		tracer:    otel.Tracer(config.ServiceName, trace.WithInstrumentationVersion(config.ServiceVersion)),
		validator: validator.New(),
	}

	return nil
}

func GormDB() *gorm.DB {
	return appCtx.gorm
}

func Config() *Configuration {
	return appCtx.config
}

func Tracer() trace.Tracer {
	return appCtx.tracer
}

func Validator() *validator.Validate {
	return appCtx.validator
}
