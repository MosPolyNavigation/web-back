package app

type App struct {
}

func New() *App {
	return &App{}
}

func (app *App) Run() error {
	return nil
}

func (app *App) Stop(done chan struct{}) error {
	return nil
}
