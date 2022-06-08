package main

//low lvl module
type InMemoryRepository struct {
	people []string
}

func NewInMemoryRepository() *InMemoryRepository {
	im := InMemoryRepository{}
	im.people = append(im.people, "Chris", "Ada")

	return &im
}

func (imr *InMemoryRepository) SearchPeople(name string) string {
	for _, p := range imr.people {
		if p == name {
			return name + " was found in repository"
		}
	}
	return "No data"
}

type SearchEngineAbstraction interface {
	FindAllPeople(name string) []string
}

type Repository struct {
	people []string
}

func (rs Repository) FindAllPeople(name string) []string {
	result := make([]string, 0)

	for _, v := range rs.people {
		if v == name {
			result = append(result, v)
		}
	}

	return result
}

//high lvl module
type SearchEngine struct {
	// breaking DIP because we deepand on concreed implementaion from low level module
	peopleInMemoryDb InMemoryRepository

	// to keep up with DIP we should take abstraction over searching people so even when implementation inside change we should keep program working
	search SearchEngineAbstraction
}

func mainDIP() {
	se := SearchEngine{}
	se.peopleInMemoryDb = *NewInMemoryRepository()
	se.peopleInMemoryDb.SearchPeople("Ada")

	// son now when underlying implementation from low lvl component change we still keeping same functionality
	se.search = Repository{}
	se.search.FindAllPeople("Ada")

}
