package dm

func (this *session[T]) FindBean(rowsSlicePtr interface{}) error {
	this.Table(rowsSlicePtr)

	return this._autoClose(func() error {
		return this.engine.sess.SQL(this.sfs.SQL(false, true), this.sfs.sqlArgs...).Find(rowsSlicePtr)
	})
}

func (this *session[T]) Find() ([]T, error) {
	vs := make([]T, 0)
	err := this._autoClose(func() error {
		return this.engine.sess.SQL(this.sfs.SQL(false, true), this.sfs.sqlArgs...).Find(&vs)
	})
	return vs, err
}

func (this *session[T]) GetBean(bean interface{}) (bool, error) {
	this.Table(bean)

	ok := false
	var err error
	err = this._autoClose(func() error {
		this.sfs.Limit(1)
		ok, err = this.engine.sess.SQL(this.sfs.SQL(true, true), this.sfs.sqlArgs...).Get(bean)
		return err
	})
	if err != nil {
		return false, err
	}
	if !ok {
		return false, err
	}
	return true, nil
}

func (this *session[T]) GetOne() (T, error) {
	v, err := this.Get()
	if err != nil {
		return *new(T), nil
	}
	return *v, nil
}

func (this *session[T]) Get() (*T, error) {
	vs := new(T)
	ok := false

	var err error
	err = this._autoClose(func() error {
		this.sfs.Limit(1)
		ok, err = this.engine.sess.SQL(this.sfs.SQL(true, true), this.sfs.sqlArgs...).Get(vs)
		return err
	})
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, err
	}
	return vs, nil
}

func (this *session[T]) Exists() (v bool, err error) {
	err = this._autoClose(func() error {
		this.sfs.Limit(1)
		v, err = this.engine.sess.SQL(this.sfs.SQL(false, true, "1"), this.sfs.sqlArgs...).Exist()
		return err
	})
	return
}

func (this *session[T]) Count() (total int64, err error) {
	err = this._autoClose(func() error {
		total, err = this.engine.sess.SQL(this.sfs.SQL(false, true, "count(1)"), this.sfs.sqlArgs...).Count()
		return err
	})
	return
}

// Page 分页查询
// size < 0 查询所有
// size = 0 只查询所有数据的量，不查询具体数据
// count = true 分页查询数据并且查询数据总量
func (this *session[T]) Page(index, size int, count bool) (vs []T, v int, err error) {
	err = this._autoClose(func() error {
		if size < 0 {
			err = this.engine.sess.SQL(this.sfs.SQL(true, true), this.sfs.sqlArgs...).Find(&vs)
			return err
		} else if size == 0 {
			_, err = this.engine.sess.SQL(this.sfs.SQL(false, true, "count(1)"), this.sfs.sqlArgs...).Get(&v)
			return err
		} else {
			if count {
				_, err = this.engine.sess.SQL(this.sfs.SQL(false, true, "count(1)"), this.sfs.sqlArgs...).Get(&v)
				if err != nil {
					return err
				}
			}

			this.sfs.Limit(size).Skip((index - 1) * size)
			err = this.engine.sess.SQL(this.sfs.SQL(true, true), this.sfs.sqlArgs...).Find(&vs)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return
}
