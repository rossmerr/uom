package mass

import "github.com/RossMerr/um"

type Mass interface {
	Unit() um.Unit
}

type Gram int64