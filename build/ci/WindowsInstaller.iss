; -- 64Bit.iss --
; Demonstrates installation of a program built for the x64 (a.k.a. AMD64)
; architecture.
; To successfully run this installation and the program it installs,
; you must have a "x64" edition of Windows.

; SEE THE DOCUMENTATION FOR DETAILS ON CREATING .ISS SCRIPT FILES!
#define SourcePath ".\app"
#define AppName "Vectorworks Utility"
#define AppExeName "Vectroworks Utility.exe"
#define AppIconName "VectorWorks.ico"

[Setup]
AppName=Vectorworks Utility
AppVersion=0.1.0
WizardStyle=modern
DefaultDirName={autopf}\Vectorworks Utility
DefaultGroupName=Vectorworks Utility
UninstallDisplayIcon={app}\Uninstall.exe
Compression=lzma2
SolidCompression=yes
OutputDir=..\package\
; "ArchitecturesAllowed=x64" specifies that Setup cannot run on
; anything but x64.
ArchitecturesAllowed=x64
; "ArchitecturesInstallIn64BitMode=x64" requests that the install be
; done in "64-bit mode" on x64, meaning it should use the native
; 64-bit Program Files directory and the 64-bit view of the registry.
ArchitecturesInstallIn64BitMode=x64

[Files]
Source: "app\VectorworksUtility.exe"; DestDir: "{app}"; DestName: "{#AppExeName}"
Source: "app\static\*"; DestDir: "{app}\static\"; Flags: recursesubdirs
Source: "app\web\*"; DestDir: "{app}\web\"; Flags: recursesubdirs
Source: "app\*.dll"; DestDir: "{app}";
Source: "assets\*.ico"; DestDir: "{app}";

[Tasks]
Name: "desktopicon"; Description: "{cm:CreateDesktopIcon}"; \
    GroupDescription: "{cm:AdditionalIcons}"; Flags: unchecked

[Icons]
Name: "{group}\Vectorworks Utility"; Filename: "{app}\{#AppExeName}"
Name: "{userdesktop}\{#AppName}"; Filename: "{app}\{#AppExeName}"; \
    IconFilename: "{app}\{#AppIconName}"; Tasks: desktopicon