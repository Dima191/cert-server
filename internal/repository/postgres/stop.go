package postgresrepository

func (r *rep) Close() {
	r.pool.Close()
}
