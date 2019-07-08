{{define "navbar"}}
<div>
	<nav class="navbar navbar-expand-lg navbar-dark bg-primary">
		<a class="navbar-brand" href="/">My Blog</a>
		<div class="collapse navbar-collapse" id="navbarColor02">
			<ul class="navbar-nav mr-auto">
				<li class="nav-item">
					<a class="nav-link admin_access" href="/category">Category <span class="sr-only">(current)</span></a>
				</li>
				<li class="nav-item">
					<a class="nav-link default_access" href="/article">Article <span class="sr-only">(current)</span></a>
				</li>
				<li class="nav-item">
					<a class="nav-link default_access" href="/photo">Photo <span class="sr-only">(current)</span></a>
				</li>
				<li class="nav-item">
					<a class="nav-link default_access" href="/about">About <span class="sr-only">(current)</span></a>
				</li>
			</ul>
			<form class="form-inline">
			
			<div class="dropdown">
				{{if .IsLogin}}
					<button  class="btn btn-outline-light" disabled>Welcome! {{.username}}</button>
					<button class="btn btn-outline-light dropdown-toggle" type="button" data-toggle="dropdown">User Log Out</button>
				{{else}}
					<button class="btn btn-outline-light dropdown-toggle" type="button" data-toggle="dropdown">User Log In</button>
				{{end}}
				<div class="dropdown-menu" aria-labelledby="dropdownMenuButton">
				{{if .IsLogin}}
					<a class="dropdown-item" href="/login">Another New User?</a>
					<a class="dropdown-item" href="#">User Profile</a>
					<a class="dropdown-item" href="/?exit=true">Log Out</a>
				{{else}}
    				<a class="dropdown-item" id="btnlogin">Log In</a>
    				<a class="dropdown-item" href="/login">New User?</a>
  				</div>
				{{end}}
			</div>
			
			</form>
		</div>
	</nav>
</div>
<!--
<script type="text/javascript">
		function redtologin() {
			window.location.href="/login"
			return false
		}
		function resetcookies() {
			window.location.href="/login?exit=true"
			return false
		}
</script>
-->

<script>
		$(document).ready(function () {
			$(".default_access",this).click(function () {
				//alert({{.IsLogin}});
				if ({{.IsLogin}}) {
				} else {
					$(this).removeAttr("href");
					$('#LoginModal').modal('show');
					//$(this).attr("href","");
				}
			});
		});
		$(document).ready(function () {
			$(".admin_access",this).click(function () {
				//alert({{.IsLogin}});
				if ({{.IsLogin}}) {
					if ({{.IsAdmin}}) {
					} else {
						alert("Only Admin Type User Can enter this page");
						$(this).attr("href","/")
					}
				} else {
					$(this).removeAttr("href");
					$('#LoginModal').modal('show');
					//$(this).attr("href","");
				}
			});
		});
		$(document).ready(function () {
			$("#btnlogin",this).click(function () {
				$(this).removeAttr("href");
				$('#LoginModal').modal('show');
				//$(this).attr("href","");
			});
		});
	</script>
{{end}}