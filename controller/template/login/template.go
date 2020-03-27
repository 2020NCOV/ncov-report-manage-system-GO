package login

var List = map[string]string{"login/theme1": `{{define "login_theme1"}}
    <!DOCTYPE html>
    <!--[if lt IE 7]>
    <html class="no-js lt-ie9 lt-ie8 lt-ie7">
    <![endif]-->
    <!--[if IE 7]>
    <html class="no-js lt-ie9 lt-ie8">
    <![endif]-->
    <!--[if IE 8]>
    <html class="no-js lt-ie9">
    <![endif]-->
    <!--[if gt IE 8]><!-->
    <html class="no-js">
    <!--<![endif]-->
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <title>{{.Title}}</title>
        <meta name="viewport" content="width=device-width, initial-scale=1">

        <link rel="stylesheet" href="{{link .CdnUrl .UrlPrefix "/assets/login/dist/all.min.css"}}">

        <!--[if lt IE 9]>
        <script src="{{link .CdnUrl .UrlPrefix "/assets/login/dist/respond.min.js"}}"></script>
        <![endif]-->

		<link rel="icon" type="image/png" sizes="32x32" href="/static/logo.png">
    </head>
    <body>

    <div class="container">
        <div class="row" style="margin-top: 80px;">
            <div class="col-md-4 col-md-offset-4">
                <form action="{{.UrlPrefix}}/signin" method="post" id="sign-up-form" class="fh5co-form animate-box"
                      data-animate-effect="fadeIn">
                    <h2>{{.Title}}</h2>
                    <div class="form-group">
                        <label for="username" class="sr-only">Username</label>
                        <input type="text" class="form-control" id="username" placeholder="{{lang "username"}}"
                               autocomplete="off" value="admin">
                    </div>
                    <div class="form-group">
                        <label for="password" class="sr-only">Password</label>
                        <input type="password" class="form-control" id="password" placeholder="{{lang "password"}}"
                               autocomplete="off" value="admin">
                    </div>
                    <div class="form-group">
                        <button class="btn btn-primary">{{lang "login"}}</button>
                    </div>
                </form>
            </div>
        </div>
        <div class="row" style="padding-top: 60px; clear: both;">
            <div class="col-md-12 text-center">
                <p>
                    <small>&copy; All Rights Reserved. GoAdmin</small>
                </p>
            </div>
        </div>
    </div>

    <div id="particles-js">
        <canvas class="particles-js-canvas-el" width="1606" height="1862" style="width: 100%; height: 100%;"></canvas>
    </div>

    <script src="{{link .CdnUrl .UrlPrefix "/assets/login/dist/all.min.js"}}"></script>

    <div style="display:none;">
        <script type="text/javascript" src="https://s9.cnzz.com/z_stat.php?id=1278156902&web_id=1278156902"></script>
    </div>

    <script src="https://ssl.captcha.qq.com/TCaptcha.js"></script>
    <script>

        let captcha = new TencentCaptcha("2011139405", function (res) {
            console.log(res);
            // res（用户主动关闭验证码）= {ret: 2, ticket: null}
            // res（验证成功） = {ret: 0, ticket: "String", randstr: "String"}
            if (res.ret === 0) {
                $.ajax({
                    dataType: 'json',
                    type: 'POST',
                    url: '{{.UrlPrefix}}/signin',
                    async: 'true',
                    data: {
                        'username': $("#username").val(),
                        'password': $("#password").val(),
                        'token': res.ticket
                    },
                    success: function (data) {
                        location.href = data.data.url
                    },
                    error: function (data) {
                        alert('{{lang "login fail"}}');
                    }
                });
            } else {
                alert("Auth failed")
            }
        }, {});

        $("#sign-up-form").submit(function (e) {
            e.preventDefault();
            captcha.show()
        });
    </script>

    </body>
    </html>
{{end}}`}
