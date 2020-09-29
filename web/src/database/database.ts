import mongoose from "mongoose";




const db =  mongoose.connect("mongodb://localhost/testDB", {
    useNewUrlParser: true,
    useFindAndModify: true,
    useUnifiedTopology: true,
    useCreateIndex: true,
  });


// 账户的数据库模型
var UserSchema = new mongoose.Schema({
    userName:String,
    passWord:String,
   
});
var User = mongoose.model('User',UserSchema);

// 新增数据
var user = {
  username: 'ydj',
  password: '123123',
  email: ''
}
var newUser = new User(user);
newUser.save();

