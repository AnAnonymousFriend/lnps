import LogoController from './controller/logo-controller';

export default [
  {
    path: '/hello',
    method: 'get',
    action: LogoController.hello
  },
  {
    path: '/login',
    method: 'post',
    action: LogoController.Login
  }
];