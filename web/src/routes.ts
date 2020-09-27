import LogoController from './controller/logo-controller';

export default [
  {
    path: '/',
    method: 'get',
    action: LogoController.hello
  }
];