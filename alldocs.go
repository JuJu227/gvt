// DO NOT EDIT THIS FILE.
//go:generate gvt help documentation

/*
gvt, a simple go vendoring tool based on gb-vendor.

Usage:
        gvt command [arguments]

The commands are:

        fetch       fetch a remote dependency
        restore     restore dependencies from manifest
        update      update a local dependency
        list        list dependencies one per line
        delete      delete a local dependency

Use "gvt help [command]" for more information about a command.


Fetch a remote dependency

Usage:
        gvt fetch [-branch branch] [-revision rev | -tag tag] [-precaire] [-no-recurse] importpath

fetch vendors an upstream import path.

The import path may include a url scheme. This may be useful when fetching dependencies
from private repositories that cannot be probed.

Flags:
	-branch branch
		fetch from the named branch. Will also be used by gvt update.
		If not supplied the default upstream branch will be used.
	-no-recurse
		do not fetch recursively.
	-tag tag
		fetch the specified tag.
	-revision rev
		fetch the specific revision from the branch or repository.
		If no revision supplied, the latest available will be fetched.
	-precaire
		allow the use of insecure protocols.

Restore dependencies from manifest

Usage:
        gvt restore [-precaire] [-connections N]

restore fetches the dependencies listed in the manifest.

It's meant for workflows that don't include checking in to VCS the vendored
source, for example if .gitignore includes lines like

    vendor/**
    !vendor/manifest

Note that such a setup requires "gvt restore" to build the source, relies on
the availability of the dependencies repositories and breaks "go get".

Flags:
	-precaire
		allow the use of insecure protocols.
	-connections
		count of parallel download connections.

Update a local dependency

Usage:
        gvt update [ -all | importpath ]

update replaces the source with the latest available from the head of the fetched branch.

Updating from one copy of a dependency to another is ONLY possible when the
dependency was fetched by branch, without using -tag or -revision. It will be
updated to the HEAD of that branch, switching branches is not supported.

To update across branches, or from one tag/revision to another, you must first
use delete to remove the dependency, then fetch [ -tag | -revision | -branch ]
to replace it.

Flags:
	-all
		update all dependencies in the manifest.
	-precaire
		allow the use of insecure protocols.

List dependencies one per line

Usage:
        gvt list [-f format]

list formats the contents of the manifest file.

Flags:
	-f
		controls the template used for printing each manifest entry. If not supplied
		the default value is "{{.Importpath}}\t{{.Repository}}{{.Path}}\t{{.Branch}}\t{{.Revision}}"

Delete a local dependency

Usage:
        gvt delete [-all] importpath

delete removes a dependency from the vendor directory and the manifest

Flags:
	-all
		remove all dependencies

*/
package main
