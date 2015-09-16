package owfs

type OwfsClient struct {
	connString string
}

func NewClient(connString string) (OwfsClient, error) {
	oc := OwfsClient{
		connString: connString,
	}
	return oc, nil
}
