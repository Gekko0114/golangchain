package lib

type Runnable interface {
	Invoke(any) (any, error)
}

type Pipeline struct {
	Pipelines []Runnable
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) Pipe(item Runnable) *Pipeline {
	p.Pipelines = append(p.Pipelines, item)
	return p
}

func (p *Pipeline) Invoke(input any) (any, error) {
	var output any
	var err error

	for _, i := range p.Pipelines {
		output, err = i.Invoke(input)
		if err != nil {
			return nil, err
		}
		input = output
	}
	return output, nil
}
