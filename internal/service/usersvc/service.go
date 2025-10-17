package usersvc

import port "meek/internal/port/service/usersvc"

// Service implements port.Service interface
type Service struct {
	// TODO: Add repository dependencies here
	// userRepo userrepo.Repository
}

// New creates a new user service
func New() port.Service {
	return &Service{
		// TODO: Inject dependencies
	}
}
