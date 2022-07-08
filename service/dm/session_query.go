package dm

func (this *session[T]) Find() ([]T, error) {
	vs := make([]T, 0)
	err := this._autoClose(func() error {
		return this.engine.sess.SQL(this.sfs.SQL(false, true), this.sfs.sqlArgs...).Find(&vs)
	})
	return vs, err
}

func (this *session[T]) Get() (T, bool, error) {
	vs := new(T)
	ok := false

	var err error
	err = this._autoClose(func() error {
		this.sfs.Limit(1)
		ok, err = this.engine.sess.SQL(this.sfs.SQL(true, true), this.sfs.sqlArgs...).Get(vs)
		return err
	})
	if err != nil {
		return *vs, false, err
	}
	if !ok {
		return *vs, false, err
	}
	return *vs, true, nil
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
