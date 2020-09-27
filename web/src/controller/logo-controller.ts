import LogoService from '../service/logo';

class HomeController {
  _service:LogoService = new LogoService()

  hello =  async(ctx) => {
    ctx.body = await this._service.hello();
  }
}

export default new HomeController()