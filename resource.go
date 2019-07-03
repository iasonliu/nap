package nap

type RestResource struct {
	Endpoint string // /get/{.userid} user=tom
	Method   string // GET
	Router   *CBRouter
}
