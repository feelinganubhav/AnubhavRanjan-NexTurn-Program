// Karma configuration
// Generated on Thu Jan 23 2025 09:28:23 GMT+0530 (India Standard Time)
const webpackConfig = {
  mode: 'development',
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: ['@babel/preset-env']
          }
        }
      }
    ]
  }
};

module.exports = function(config) {
  config.set({

    // base path that will be used to resolve all patterns (eg. files, exclude)
    basePath: '',


    // frameworks to use
    // available frameworks: https://www.npmjs.com/search?q=keywords:karma-adapter
    frameworks: ['jasmine', 'webpack'],


    // list of files / patterns to load in the browser
    files: [
      'test/**/*.spec.js',
      'src/**/*.js'
    ],


    // list of files / patterns to exclude
    exclude: [
    ],

    webpack: webpackConfig, // Pass Webpack configuration here

    // preprocess matching files before serving them to the browser
    // available preprocessors: https://www.npmjs.com/search?q=keywords:karma-preprocessor
    preprocessors: {
      'test/**/*.spec.js': ['webpack'], // Use webpack to process test files
      'src/**/*.js': ['webpack'],      // Use webpack to process source files
    },

    

    // test results reporter to use
    // possible values: 'dots', 'progress'
    // available reporters: https://www.npmjs.com/search?q=keywords:karma-reporter
    reporters: ['progress', 'kjhtml', 'coverage', 'junit'],

    coverageReporter: {
      type: 'html', // Output HTML report
      dir: 'coverage/' // Output directory
    },
    
    junitReporter: {
      outputDir: 'test-results', // Save XML reports here
      outputFile: 'results.xml', // File name
      useBrowserName: false      // Avoid appending browser name to file
    },
    


    // web server port
    port: 9876,


    // enable / disable colors in the output (reporters and logs)
    colors: true,


    // level of logging
    // possible values: config.LOG_DISABLE || config.LOG_ERROR || config.LOG_WARN || config.LOG_INFO || config.LOG_DEBUG
    logLevel: config.LOG_INFO,


    // enable / disable watching file and executing tests whenever any file changes
    autoWatch: true,


    // start these browsers
    // available browser launchers: https://www.npmjs.com/search?q=keywords:karma-launcher
    browsers: ['Chrome'],


    // Continuous Integration mode
    // if true, Karma captures browsers, runs the tests and exits
    singleRun: false,

    // Concurrency level
    // how many browser instances should be started simultaneously
    concurrency: Infinity
  })
}
