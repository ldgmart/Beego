<html>
<head>
<title></title>
</head>
<body>

  <script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"></script>

 
<div id="app">
  {{ message }}
</div>

<script> 
var app = new Vue({
  el: '#app',
  data: {
    message: 'Hello Vue!'
  }
})</script>

    <div class="container">
        <div class="row vertical-offset-75">
          <div class="col-md-6 col-md-offset-3">
            <div class="panel panel-default">
              <div class="panel-heading text-center">
                <h3 class="panel-title"><strong>Login Page </strong></h3>
             </div> 

              <div class="panel-body">
                <form accept-charset="utf-8" role="form" class="form-horizontal" method="POST" action='/Login'>

                          <div class="form-group">
                           <div class="col-sm-9">
                              ID : <input class="form-control" name="id"  type="text" />
                            </div>
                          </div>
                          <div class="form-group">
                           <div class="col-sm-9">
                    PW : <input class="form-control" name="password" type="password" />
                    <input class="btn btn-lg btn-success" type="submit" value="Login"> 
                            </div>
           
                          </div>
   
                        </form>
              </div>
    
                    <div class="panel-footer text-center clearfix"> <a href='/Signup'>Signup</a></div>
    
          </div>
        </div>
      </div>
    </div>
</body>
</html>
<!--
<form action ="/Login" method="post">
   
   Username:<input type="text" name ="id">
   Password:<input type="password" name ="password">
   <input type="submit" value="Login" >
 </form>
<button type="button" onclick="location.href='/Signup' ">Signup</button> 
 </body>
 </html>
-->
