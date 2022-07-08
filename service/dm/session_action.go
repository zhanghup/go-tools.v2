package dm

func (this *session[T]) Order(order ...string) ISession[T] {
	this.sfs.orderby = append(this.sfs.orderby, order...)
	return this
}

func (this *session[T]) SF(sqlstr string, querys ...any) ISession[T] {
	this.sfs.SF(sqlstr, querys...)
	return this
}
