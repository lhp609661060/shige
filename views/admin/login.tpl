<form class="form-signin" method="post">
    <label for="username" class="sr-only">用户名</label>
    <input type="text" name="username" id="username" class="form-control" placeholder="用户名" required autofocus>
    {{if eq .InputName "username"}}<span class="error">{{.ErrMsg}}</span>{{end}}
    <br>
    <label for="password" class="sr-only">密码</label>
    <input type="password" name="password" id="password" class="form-control" placeholder="密码" required>
    {{if eq .InputName "password"}}<span class="error">{{.ErrMsg}}</span>{{end}}
    <div class="checkbox">
        <label>
            <input type="checkbox" value="remember-me"> 记住登录
        </label>
    </div>
    <button class="btn btn-lg btn-primary btn-block" type="submit">登录</button>
</form>
