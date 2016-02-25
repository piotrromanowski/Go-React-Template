module.exports = {
  entry: [
    './frontend/main.js'
  ],
  output: {
    path: __dirname,
    filename: '/backend/public/bundle.js'
  },
  module: {
    loaders: [{
      exclude: /node_modules/,
      loader: 'babel'
    }]
  },
  resolve: {
    extensions: ['', '.js', '.jsx']
  },
  devServer: {
    historyApiFallback: true,
    contentBase: './'
  }
};
