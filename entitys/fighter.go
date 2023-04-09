package entitys

type Fighter struct {
	Stats
	Moveset []Movement
}

type Stats struct {
	Life, Speed, Normaldmg, Specialdmg, Normaldfs, Specialdfs int
}
