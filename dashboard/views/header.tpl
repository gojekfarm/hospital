{{ define "header" }}

<!doctype html>
<html lang="en">
  <head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
    <title>Hospital</title>
  </head>
  <body>
        <nav class="navbar navbar-expand-lg navbar-light bg-light justify-content-between">
                <a class="navbar-brand" href="#">Hospital</a>
                <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
                  <span class="navbar-toggler-icon"></span>
                </button>
                <div class="collapse navbar-collapse" id="navbarText">
                  <ul class="navbar-nav">
                    <li class="nav-item">
                        {{if eq .Page "mapping"}}
                          <a class="nav-link active" href="/dashboard">Mapping</a>
                        {{else}}
                          <a class="nav-link" href="/dashboard">Mapping</a>
                        {{end}}
                    </li>
                    <li class="nav-item ">
                        {{if eq .Page "logs"}}
                            <a class="nav-link active" href="/dashboard/logs">Logs</a>
                        {{else}}
                          <a class="nav-link" href="/dashboard/logs">Logs</a>
                        {{end}}
                    </li>
                    <li class="nav-item ">
                        {{if eq .Page "summary"}}
                            <a class="nav-link active" href="/dashboard/summary">Logs</a>
                        {{else}}
                          <a class="nav-link" href="/dashboard/summary">Logs</a>
                        {{end}}
                    </li>
                  </ul>
                  
                </div>
                {{if eq .Page "mapping"}}
                  <span class="navbar-text">
                        <button type="button" class="btn btn-outline-success my-2 my-sm-0" data-toggle="modal" data-target="#myModal">Add Mapping</button>
                  </span>
                {{end}}
              </nav>

{{ end }}