import LogoService from '../service/logo';
import LoginService from '../service/Login'
import { Console } from 'console';

class HomeController {
  _service:LogoService = new LogoService()
  _loginService:LoginService = new LoginService()

  hello =  async(ctx) => {
    ctx.body = await this._service.hello();
  }

  Login = async(ctx) => {
    ctx.body = ctx.request.body;
    console.log("用户信息" + ctx.request.body.UserName);
    //ctx.body = await this._loginService.Login(ctx.body)
  }



}

export default new HomeController()