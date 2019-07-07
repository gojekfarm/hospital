{{ template "header" .}}

<div class="container">
        <br>
        <div class="row">
        <div class="col-md-12"><table class="table table-hover">
                <thead class="thead-light">
                    <tr>
                    <th scope="col">Application ID</th>
                    <th scope="col"><i class="fa fa-check-circle" style="color:green"></i> Resolved Alerts</th>
                    <th scope="col"><i class="fa fa-times-circle" style="color:red"></i> Failed Alerts</th>
                    <th scope="col"><i class="fa fa-free-code-camp" style="color:yellow"></i> Firing Alerts</th>
                    </tr>
                </thead>
                <tbody>
                    {{with .Summary}}
                    <tr>
                    <th>{{.ApplicationID}}</th>
                    <td>{{.Success}}</td>
                    <td>{{.Fail}}</td>
                    <td>{{.Firing}}</td>
                    </tr>
                    {{end}}
                </tbody>
            </table>
        </div>
    </div>
</div>

<div class="container">
        <br>
        <table class="table table-dark">
            <thead>
            <tr>
                <th scope="col">Script</th>
                <th scope="col">Status</th>
                <th scope="col">Logs</th>
            </tr>
            </thead>
            <tbody>
                {{range .Logs}}
                    <tr>
                        <td>{{.Script}}</td>
                        {{if eq .Status "completed"}}
                            <td><i class="fa fa-check-circle" style="color:green"></i> {{.Status}}</td>
                        {{else if eq .Status "failed"}}
                            <td><i class="fa fa-times-circle" style="color:red"></i> {{.Status}}</td>
                        {{else}}
                            <td><i class="fa fa-free-code-camp" style="color:yellow"></i> {{.Status}}</td>
                        {{end}}
                        <td><a href="/dashboard/logs/{{.ID}}">{{.Logs}}</a></td>
                    </tr>
                {{end}}
            </tbody>
        </table>
</div>

{{ template "footer" .}}