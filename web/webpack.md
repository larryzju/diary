## Overview

compile Javascript modules


### package.json

在 package.json 中定义 "scripts" 字段（相当于 shell 中的别名），如

```json
{
    "build": "webpack -d"
}
```

可以通过 `yarn run build` 来调用 `webpack -d`.

webpack -d 参数表示保留 debug 信息，可以方便开发 (product 模式下代码将被压缩)



### asset management

通过 loader rules （匹配文件名，使用不同的 load，在代码中用 import 语法引入
css, image, file 等文件），更有效的管理 asset （不需要把所有的文件放在同一个 /assets 目录）

```json
module: {
    rules: [
	{
	    test: /\.css$/,
	    use: [
		'style-loader',
		'css-loader',
	    ],
	},
	{
	    test: /\.(png|svg|jpg|gif)$/,
	    use: [
		'file-loader',
	    ],
	},
	{
	    test: /\.(woff|woff2|eot|ttf|otf)$/,
	    use: [
		'file-loader',
	    ],
	},
	{
	    test: /\.(csv|tsv)$/,
	    use: [
		'csv-loader',
	    ],
	},
	{
	    test: /\.xml$/,
	    use: [
		'xml-loader',
	    ],
	},
	{
	    test: /\.yaml$/,
	    type: 'json',
	    use: [
		'yaml-loader',
	    ],
	},
    ],
}
```




### Output Management

webpack.config.js 中可以有多个 entry，分别有自己的别名，输出时可以用 `[name].bundle.js` 作为模块，自动带入别名

```json
module.exports = {
    entry: {
        app: './src/index.js',
        print: './src/print.js',
    },
    output: {
        filename: '[name].bundle.js',
        path: path.resolve(__dirname, 'dist'),
    }
}
```

可以使用插件自动生成 html 页面，如
* html-webpack-plugin
* html-webpack-template

webpack 维护了一份 manifest，记录了 module 到 bundle 的对应生成关系，可以用 webpack-manifest-plugin 导出

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




## Development

### source map

使用 source map，为生成的 bundle.js 添加到源文件的映射信息，方法调试。在 webpack.config.js 中添加

```json
{
    mode: "development",
    devtool: "inline-source-map"
}
```

### watch mode

类似 entr 功能，webpack --watch 参数

### webpack-dev-server


## Plugins

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

plugins array 中写入 plugin instance，常用的插件如

| plugin                  | description                  |
|-------------------------|------------------------------|
| clean-webpack-plugin    | 清理 dist 目录（放在第一位） |
| html-webpack-plugin     | 自动生成 html 页面           |
| html-webpack-template   | 同上                         |
| webpack-manifest-plugin | 导出 manifest 到 json        |
| webpack-dev-server      | 可实现浏览器自动刷新         |
