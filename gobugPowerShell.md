Windows PowerShell
Copyright (C) Microsoft Corporation. All rights reserved.

PS C:\1Go\SmallPrograms\gobug> go mod tidy
PS C:\1Go\SmallPrograms\gobug> wails dev
Wails CLI v2.12.0

Updating go.mod to use Wails 'v2.12.0'
Executing: go mod tidy
  • Generating bindings: Done.
  • No Install command. Skipping.
  • No Build command. Skipping.

  ERROR   unable to auto discover frontend:dev:serverUrl without a frontend:dev:watcher command, please either set frontend:dev:watcher or remove the auto discovery from frontend:dev:serverUrl
 ♥   If Wails is useful to you or your company, please consider sponsoring the project:
https://github.com/sponsors/leaanthony
PS C:\1Go\SmallPrograms\gobug>