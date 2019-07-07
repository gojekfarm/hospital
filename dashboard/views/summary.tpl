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
                    {{range .Summaries}}
                    <tr>
                    <th><a href="/dashboard/summary/{{.ApplicationID}}">{{.ApplicationID}}</a></th>
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

{{ template "footer" .}}