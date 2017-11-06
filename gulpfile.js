//plug-in
var gulp = require('gulp');
var pug = require('gulp-pug');
var sass = require('gulp-sass');
var ts = require('gulp-typescript');
var sourcemaps = require('gulp-sourcemaps');
var webpack = require('webpack-stream');
var named = require('vinyl-named');
var plumber = require('gulp-plumber');
 

gulp.task('pug', () => {
    gulp.src(['./dev/pug/**/*.pug', '!./dev/pug/**/_*.pug'], { base: 'dev/pug' })
        .pipe(plumber())
        .pipe(pug({
            pretty: true
        }))
        .pipe(gulp.dest('./app/views/'));
});


gulp.task('sass', function() {
    gulp.src(['./dev/sass/**/*.sass', '!./dev/sass/**/_*.sass'],{ base: 'dev/sass' })
        .pipe(plumber())
        .pipe(sourcemaps.init())
        .pipe(sass())
        .pipe(sourcemaps.write('../map/sassMap'))
        .pipe(gulp.dest('./public/css/'));
});


var webpackConfig = require('./webpack.config.js');
 
gulp.task('ts', function () {
    gulp.src(['./dev/typeScript/**/*.ts', '!./dev/typeScript/**/_*.ts','!./node_modules/**'],{base:'dev/typeScript'})
    .pipe(plumber())
    .pipe(named(function(file) {
      return file.relative.replace(/\.[^\.]+$/, '');
    }))
    .pipe(webpack(webpackConfig))
    .pipe(gulp.dest('./public/js/'));
});


gulp.task('default', ['pug','sass','ts'], function() {
    gulp.watch(['./dev/pug/**/*.pug'], ['pug']);
    gulp.watch(['./dev/sass/**/*.sass'], ['sass']);
    gulp.watch(['./dev/typeScript/**/*.ts'], ['ts']);
})