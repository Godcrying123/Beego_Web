{{define "alert"}}
<div class="modal fade" id="AlertModal" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel" aria-hidden="true">
    <div class="modal-dialog" role="document">
        <div class="container">
            {{if .success}}
            <div class="alert alert-success alert-dismissible fade show">
                <!--<button type="button" class="close" data-dismiss="alert">&times;</button>
                -->
                <strong>Success!</strong> You should <a href="#" class="alert-link">read this message</a>.
            </div>
            {{else}} {{end}} {{if .info}}
            <div class="alert alert-info alert-dismissible fade show">
                 <!--<button type="button" class="close" data-dismiss="alert">&times;</button>
                -->
                <strong>Info!</strong> You should <a href="#" class="alert-link">read this message</a>.
            </div>{{else}}{{end}} {{if .warning}}
            <div class="alert alert-warning alert-dismissible fade show">
                 <!--<button type="button" class="close" data-dismiss="alert">&times;</button>
                -->
                <strong>Warning!</strong> You should <a href="#" class="alert-link">read this message</a>.
            </div>
            {{else}} {{end}}{{if .danger}}
            <div class="alert alert-danger alert-dismissible fade show">
                 <!--<button type="button" class="close" data-dismiss="alert">&times;</button>
                -->
                <strong>Danger!</strong> You should <a href="#" class="alert-link">read this message</a>.
            </div>
            {{else}}{{end}}
        </div>
    </div>
</div>
<script>
$(document).ready(function() {
	if ({{.success}}||{{.info}}||{{.warning}}||{{.danger}}) {
	$('#AlertModal').modal('show');}
	});
</script>
{{end}}
