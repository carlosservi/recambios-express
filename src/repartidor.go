package main

type Repartidor struct {
	id      int
	horario Horario
	pedidos Pedido
}

type Horario struct {
	inicio string
	fin    string
}

func NuevoRepartidor(id int, horario Horario) *Repartidor {
	return &Repartidor{
		id:      id,
		horario: horario,
	}
}

func (r *Repartidor) ID() int {
	return r.id
}

func (r *Repartidor) Horario() Horario {
	return r.horario
}

func (r *Repartidor) Pedidos() Pedido {
	return r.pedidos
}