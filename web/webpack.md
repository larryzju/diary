## Overview

compile Javascript modules

## Concepts

**webpack** is a static module bundler for modern Javascript applications.

- build dependency graph and 
- generates one or more bundles

| Term    | Description                                                  |
| ------- | ------------------------------------------------------------ |
| Entry   | entry point source file to start building out its internal *dependency graph* |
| Output  | output directory and filename, default file is `./dist/main.js` |
| Loaders | process other types of files then JS or JSON, convert them to modules |
| Plugins | extend webpack's capabilities: optimization, asset management, inject env... |
| Mode    | development, production (default), none                      |

### Loaders

怎样把 css, image 打包到 module?

```js
const path = require('path');

module.exports = {
  output: {
    filename: 'my-first-webpack.bundle.js'
  },
  module: {
    rules: [
      { test: /\.txt$/, use: 'raw-loader' }
    ]
  }
};
```

对 txt 文件使用 raw-loader。当 require txt 文件时，将调用 raw-loader 来加载之



### Plugins

```js
const HtmlWebpackPlugin = require('html-webpack-plugin'); //installed via npm
const webpack = require('webpack'); //to access built-in plugins

module.exports = {
  module: {
    rules: [
      { test: /\.txt$/, use: 'raw-loader' }
    ]
  },
  plugins: [
    new HtmlWebpackPlugin({template: './src/index.html'})
  ]
};
```

plugins array 中写入 plugin instance