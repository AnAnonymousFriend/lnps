const Koa = require('koa');

// 注意require('koa-router')返回的是函数:
const router = require('koa-router')();
const bodyParser = require('koa-bodyparser');

const app = new Koa();


// x-responese-time

app.use(bodyParser());
app.use(router.routes());

app.use(async(ctx,next)=>{
    console.log(`${ctx.request.method} ${ctx.request.url}`); // 打印URL
    await next(); // 调用下一个middleware
})

app.use(async(ctx,next) =>{
    const start = new Date().getTime();
    await next();
    const ms = new Date().getTime() - start;
    console.log(`Time: ${ms}ms`)
})



app.use(async(ctx,next)=>{
   if(ctx.request.path == "/test"){
       ctx.response.body = `Test Page`;
   }else{
       await next();
   }
})


app.use(async(ctx,next)=>{
    if(ctx.request.path == "/error"){
        ctx.response.body = `Error Page`;
    }else{
        await next();
    }
 })

 router.get('/hello/:name', async (ctx, next) => {
    var name = ctx.params.name;
    ctx.response.body = `<h1>Hello, ${name}!</h1>`;
});

router.post('/signin', async (ctx, next) => {
    var
        name = ctx.request.body.name || '',
        password = ctx.request.body.password || '';
    console.log(`signin with name: ${name}, password: ${password}`);
    if (name === 'koa' && password === '12345') {
        ctx.response.body = `<h1>Welcome, ${name}!</h1>`;
    } else {
        ctx.response.body = `<h1>Login failed!</h1>
        <p><a href="/">Try again</a></p>`;
    }
});




app.listen(3000);
console.log('app started at port 3000...');