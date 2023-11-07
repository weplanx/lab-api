// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package bootstrap

import (
	"github.com/weplanx/go/rest"
	"github.com/weplanx/go/sessions"
	"github.com/weplanx/go/values"
	"github.com/weplanx/server/api"
	"github.com/weplanx/server/api/acc_tasks"
	"github.com/weplanx/server/api/builders"
	"github.com/weplanx/server/api/clusters"
	"github.com/weplanx/server/api/datasets"
	"github.com/weplanx/server/api/endpoints"
	"github.com/weplanx/server/api/imessages"
	"github.com/weplanx/server/api/index"
	"github.com/weplanx/server/api/lark"
	"github.com/weplanx/server/api/monitor"
	"github.com/weplanx/server/api/projects"
	"github.com/weplanx/server/api/queues"
	"github.com/weplanx/server/api/tencent"
	"github.com/weplanx/server/api/workflows"
	"github.com/weplanx/server/common"
	"github.com/weplanx/server/openapi"
	index3 "github.com/weplanx/server/openapi/index"
	"github.com/weplanx/server/xapi"
	"github.com/weplanx/server/xapi/emqx"
	index2 "github.com/weplanx/server/xapi/index"
)

// Injectors from wire.go:

func NewAPI(values2 *common.Values) (*api.API, error) {
	client, err := UseMongoDB(values2)
	if err != nil {
		return nil, err
	}
	database := UseDatabase(values2, client)
	redisClient, err := UseRedis(values2)
	if err != nil {
		return nil, err
	}
	conn, err := UseNats(values2)
	if err != nil {
		return nil, err
	}
	jetStreamContext, err := UseJetStream(conn)
	if err != nil {
		return nil, err
	}
	keyValue, err := UseKeyValue(values2, jetStreamContext)
	if err != nil {
		return nil, err
	}
	cipher, err := UseCipher(values2)
	if err != nil {
		return nil, err
	}
	captcha := UseCaptcha(redisClient)
	locker := UseLocker(redisClient)
	clientClient, err := UseTransfer(jetStreamContext)
	if err != nil {
		return nil, err
	}
	inject := &common.Inject{
		V:         values2,
		Mgo:       client,
		Db:        database,
		RDb:       redisClient,
		Nats:      conn,
		JetStream: jetStreamContext,
		KeyValue:  keyValue,
		Cipher:    cipher,
		Captcha:   captcha,
		Locker:    locker,
		Transfer:  clientClient,
	}
	hertz, err := UseHertz(values2)
	if err != nil {
		return nil, err
	}
	csrf := UseCsrf(values2)
	service := UseValues(keyValue, cipher)
	controller := &values.Controller{
		Service: service,
	}
	sessionsService := UseSessions(values2, redisClient)
	sessionsController := &sessions.Controller{
		Service: sessionsService,
	}
	restService := UseRest(values2, client, database, redisClient, jetStreamContext, keyValue, cipher)
	restController := &rest.Controller{
		Service: restService,
	}
	passport := UseAPIPassport(values2)
	tencentService := &tencent.Service{
		Inject: inject,
	}
	indexService := &index.Service{
		Inject:   inject,
		Sessions: sessionsService,
		Passport: passport,
		TencentX: tencentService,
	}
	indexController := &index.Controller{
		V:      values2,
		Csrf:   csrf,
		IndexX: indexService,
	}
	tencentController := &tencent.Controller{
		TencentX: tencentService,
	}
	larkService := &lark.Service{
		Inject:   inject,
		Sessions: sessionsService,
		Locker:   locker,
		Passport: passport,
		IndexX:   indexService,
	}
	larkController := &lark.Controller{
		V:        values2,
		Passport: passport,
		LarkX:    larkService,
		IndexX:   indexService,
	}
	clustersService := &clusters.Service{
		Inject: inject,
	}
	projectsService := &projects.Service{
		Inject:    inject,
		ClustersX: clustersService,
	}
	projectsController := &projects.Controller{
		ProjectsX: projectsService,
	}
	clustersController := &clusters.Controller{
		ClustersX: clustersService,
	}
	endpointsService := &endpoints.Service{
		Inject: inject,
	}
	endpointsController := &endpoints.Controller{
		EndpointsX: endpointsService,
	}
	workflowsService := &workflows.Service{
		Inject:     inject,
		EndpointsX: endpointsService,
	}
	workflowsController := &workflows.Controller{
		WorkflowsX: workflowsService,
	}
	queuesService := &queues.Service{
		Inject:    inject,
		ProjectsX: projectsService,
	}
	queuesController := &queues.Controller{
		QueuesX: queuesService,
	}
	imessagesService := &imessages.Service{
		Inject: inject,
	}
	imessagesController := &imessages.Controller{
		ImessagesX: imessagesService,
	}
	acc_tasksService := &acc_tasks.Service{
		Inject:   inject,
		TencentX: tencentService,
	}
	acc_tasksController := &acc_tasks.Controller{
		AccTasksX: acc_tasksService,
	}
	datasetsService := &datasets.Service{
		Inject: inject,
		Values: service,
	}
	datasetsController := &datasets.Controller{
		DatasetsX: datasetsService,
	}
	influxdb2Client := UseInflux(values2)
	monitorService := &monitor.Service{
		Inject: inject,
		Flux:   influxdb2Client,
	}
	monitorController := &monitor.Controller{
		MonitorX: monitorService,
	}
	buildersService := &builders.Service{
		Inject: inject,
	}
	buildersController := &builders.Controller{
		BuildersX: buildersService,
	}
	apiAPI := &api.API{
		Inject:     inject,
		Hertz:      hertz,
		Csrf:       csrf,
		Values:     controller,
		Sessions:   sessionsController,
		Rest:       restController,
		Index:      indexController,
		IndexX:     indexService,
		Tencent:    tencentController,
		TencentX:   tencentService,
		Lark:       larkController,
		LarkX:      larkService,
		Projects:   projectsController,
		ProjectsX:  projectsService,
		Clusters:   clustersController,
		ClustersX:  clustersService,
		Endpoints:  endpointsController,
		EndpointsX: endpointsService,
		Workflows:  workflowsController,
		WorkflowsX: workflowsService,
		Queues:     queuesController,
		QueuesX:    queuesService,
		Imessages:  imessagesController,
		ImessagesX: imessagesService,
		AccTasks:   acc_tasksController,
		AccTasksX:  acc_tasksService,
		Datasets:   datasetsController,
		DatasetsX:  datasetsService,
		Monitor:    monitorController,
		MonitorX:   monitorService,
		Builders:   buildersController,
		BuildersX:  buildersService,
	}
	return apiAPI, nil
}

func NewXAPI(values2 *common.Values) (*xapi.API, error) {
	client, err := UseMongoDB(values2)
	if err != nil {
		return nil, err
	}
	database := UseDatabase(values2, client)
	redisClient, err := UseRedis(values2)
	if err != nil {
		return nil, err
	}
	conn, err := UseNats(values2)
	if err != nil {
		return nil, err
	}
	jetStreamContext, err := UseJetStream(conn)
	if err != nil {
		return nil, err
	}
	keyValue, err := UseKeyValue(values2, jetStreamContext)
	if err != nil {
		return nil, err
	}
	cipher, err := UseCipher(values2)
	if err != nil {
		return nil, err
	}
	captcha := UseCaptcha(redisClient)
	locker := UseLocker(redisClient)
	clientClient, err := UseTransfer(jetStreamContext)
	if err != nil {
		return nil, err
	}
	inject := &common.Inject{
		V:         values2,
		Mgo:       client,
		Db:        database,
		RDb:       redisClient,
		Nats:      conn,
		JetStream: jetStreamContext,
		KeyValue:  keyValue,
		Cipher:    cipher,
		Captcha:   captcha,
		Locker:    locker,
		Transfer:  clientClient,
	}
	hertz, err := UseHertz(values2)
	if err != nil {
		return nil, err
	}
	service := &index2.Service{
		Inject: inject,
	}
	controller := &index2.Controller{
		IndexService: service,
	}
	emqxService := &emqx.Service{
		Inject: inject,
	}
	emqxController := &emqx.Controller{
		EmqxService: emqxService,
	}
	xapiAPI := &xapi.API{
		Inject:       inject,
		Hertz:        hertz,
		Index:        controller,
		IndexService: service,
		Emqx:         emqxController,
		EmqxService:  emqxService,
	}
	return xapiAPI, nil
}

func NewOpenAPI(values2 *common.Values) (*openapi.API, error) {
	client, err := UseMongoDB(values2)
	if err != nil {
		return nil, err
	}
	database := UseDatabase(values2, client)
	redisClient, err := UseRedis(values2)
	if err != nil {
		return nil, err
	}
	conn, err := UseNats(values2)
	if err != nil {
		return nil, err
	}
	jetStreamContext, err := UseJetStream(conn)
	if err != nil {
		return nil, err
	}
	keyValue, err := UseKeyValue(values2, jetStreamContext)
	if err != nil {
		return nil, err
	}
	cipher, err := UseCipher(values2)
	if err != nil {
		return nil, err
	}
	captcha := UseCaptcha(redisClient)
	locker := UseLocker(redisClient)
	clientClient, err := UseTransfer(jetStreamContext)
	if err != nil {
		return nil, err
	}
	inject := &common.Inject{
		V:         values2,
		Mgo:       client,
		Db:        database,
		RDb:       redisClient,
		Nats:      conn,
		JetStream: jetStreamContext,
		KeyValue:  keyValue,
		Cipher:    cipher,
		Captcha:   captcha,
		Locker:    locker,
		Transfer:  clientClient,
	}
	hertz, err := UseHertz(values2)
	if err != nil {
		return nil, err
	}
	service := &index3.Service{
		Inject: inject,
	}
	controller := &index3.Controller{
		IndexService: service,
	}
	openapiAPI := &openapi.API{
		Inject:       inject,
		Hertz:        hertz,
		Index:        controller,
		IndexService: service,
	}
	return openapiAPI, nil
}
