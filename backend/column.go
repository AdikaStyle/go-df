package backend

func (c *Columns) GetHeaders() Headers {
	var out Headers
	for k, v := range *c {
		out = append(out, Header{
			Name:    k,
			Kind:    v[0].Kind(),
			Visible: true,
		})
	}
	return out
}
