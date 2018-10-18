const path = require('path')

const isProd = process.env.NODE_ENV === 'production'

module.exports = {
  module: {
    rules: [{
      include: [path.resolve(__dirname, 'src')],
      loader: 'babel-loader',
      test: /\.js$/
    }]
  },

  entry: './src/index.js',

  output: {
    filename: '[name].js',
    path: path.resolve(__dirname, 'dist')
  },

  mode: isProd ? 'production' : 'development',

  optimization: {
    splitChunks: {
      cacheGroups: {
        vendors: {
          priority: -10,
          test: /[\\/]node_modules[\\/]/
        }
      },

      chunks: 'async',
      minChunks: 1,
      minSize: 30000,
      name: true
    }
  }
}
