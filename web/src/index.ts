import Koa = require('koa');
import Router = require('koa-router');
import bodyParser = require('koa-bodyparser');
import AppRoutes from './routes';
import cors = require('koa2-cors')

const app = new Koa();
const router = new Router();
const port = process.env.PORT || 3000;


//路由
AppRoutes.forEach(route => router[route.method](route.path, route.action));
app.use(cors({
    exposeHeaders: ['WWW-Authenticate', 'Server-Authorization', 'Date'],
    maxAge: 100,
    credentials: true,
    allowMethods: ['GET', 'POST', 'OPTIONS'],
    allowHeaders: ['Content-Type', 'Authorization', 'Accept', 'X-Custom-Header', 'anonymous'],
}));
app.use(bodyParser());
app.use(router.routes());
app.use(router.allowedMethods());
app.listen(port);

console.log(`应用启动成功 端口:${port}`);