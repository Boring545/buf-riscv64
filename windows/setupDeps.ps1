$installedApps = choco list --localonly | Out-String

$apps = @('diffutils', 'golang', 'protoc')
$apps | ForEach-Object {
    if (-Not $installedApps.Contains($PSItem)) {
        Write-Host "Installing $PSItem"
        choco install --confirm $PSItem
    }
}