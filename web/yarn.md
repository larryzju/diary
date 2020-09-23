## Dependencies

### Commands

- `yarn install` similar to `go get`, download all dependencies and records them in `yarn.lock` file
- `yarn add package@version`

### Types

| type                 | purposes                                              |
| -------------------- | ----------------------------------------------------- |
| dependencies         | runtime dependencies                                  |
| devDependencies      | babel, flow..., only in development workflow          |
| peerDenendencies     | ?? only come up when publishing your own package      |
| optionalDependencies |                                                       |
| bundledDependencies  | as/use local library,  packed when running`yarn pack` |

### Version & Tag

## Topics

### package.json

```json
{
    // Essential part
    // name should name have an uppercase letter in the name
    // as part of URL and directory name in `node_modules`
    "name": "package-name", 
    "version": "1.0.0",
    
    // Info
    "description": "help people understand the purpose of the package",
    "keywords": ["keyword", "for", "search", "in", "package manager"],
    "license": "MIT",
    
    // Links
    "homepage": "https://your-package.org",
    "bugs": "https://github.com/user/repo/issues",
    "repository": {
        "type": "git",
        "url": "https://github.com/larryzju/somerepo.git",
    },
    
    // Maintainers
    "author": {
        "name": "your name",
        "email": "you@example.com",
        "url": "http://your-website.com"
    },
    "contributors": {    
    },

	// Files
	"files": [], // file, directory, pattern to include files
    "main": "filename.js", // primary entry point
    "bin":  {
        "command1": "bin/command1.js",
        "command2": "bin/command2"
    },
    "man": ["./man/doc.1", "./man/doc.2"], // man pages
    "directories": {	// customize locations to put files
        "lib": "path/to/lib/",
        "bin": "path/to/bin/",
        "man": "path/to/man/",
        "doc": "path/to/doc/",
        "example": "path/to/example/""
    },
    
    // Tasks
    // include runnable scripts or other configuration
    "scripts": {  // like shell alias
        "build-project": "node build-project.js",
        // there're some predefined script name, 
        // such as install, postinstall, prepublish, prepare 
        // will be called automatically in certain lifecycle!
    },
    
    "config": {
        "port": "8080",
    },
    
    // Dependencies
    // both development and production
    "dependencies": {
        "package-1": "^3.1.4",
    },
    // only development
    "devDependencies": {
    	"package-2": "^0.4.2",
	},
    // state compatibility of you package with versions of other packages
    "peerDependencies": {
        "package-3": "^2.7.18",
    },
    // add metadata to peer dependencies
    "peerDependenciesMeta": {
        "package-3": {
            // used to supress the warning ofr a missing peer dependencies
            "optional": true
        }
    },
    // mark the package is not required
    "optionalDependencies": {
        "package-5": "^1.6.1",
    },
    // bundle together when publishing your package
    "bundledDependencies": ["package-4"],
    // only allow one version of a given dependency
    "flat": "true",
    // override a version of a particular nested dependency ??
    "resolutions": {
        "transitive-package-1": "0.0.29",
	},
    
  
    // System
    // specify versions of clients used with this package
    "engine": {
        "node": ">=4.4.7 <7.0.0",	// process.versions
        "zlib": "^1.2.8",
        "yarn": "^0.14.0",
    },
    "os": ["!win32"], // process.platform
    "cpu": ["!mips", "!arm"], // process.arch
    
    
    // Publishing
    "private": true,	// don't published in a package manager
    "publishConfig": {},
}
```

### yarn.lock

- Dependencies tree file to creates reproducible results
- Generate automatically by Yarn
- Only used in current, top-level package
- Should be checked into source control system.

### scoped package

TODO