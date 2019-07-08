{{define "modal"}}
<div class="modal fade" id="ArticleModal{{.Id}}" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel" aria-hidden="true">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title">{{.Title}}</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                                        <span aria-hidden="true">&times;</span>
                                        </button>
            </div>
            <div class="modal-body">
                <div class="card text-center">
                    <div class="card-header">
                        <ul class="nav nav-tabs card-header-tabs">
                            <li class="nav-item">
                                <a class="nav-link active" href="#">Content</a>
                            </li>
                            <li class="nav-item">
                                <a class="nav-link active" href="#">Comment</a>
                            </li>
                            <li class="nav-item">
                                <a class="nav-link active" href="#">Photos</a>
                            </li>

                        </ul>
                    </div>
                    <div class="card-body">
                        <h5 class="card-title">{{.Content}}</h5>
                        <!--
                        <p class="card-text">With supporting text below as a natural lead-in to additional content.</p>
                        -->
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-primary">Comment</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>

{{end}}