var gulp = require('gulp');
var concat = require('gulp-concat');
var uglify = require('gulp-uglify');
var reactify = require('reactify');
var browserify = require('browserify');
var streamify = require('gulp-streamify');
var mocha = require('gulp-mocha');
var babelify = require("babelify");
var babel = require("gulp-babel");
var es2015 = require('babel-preset-es2015');
var clean = require('gulp-clean');

var source = require('vinyl-source-stream');


var path = {
    HTML: 'app/index.html',
    JS: ['app/js/src/*.js'],
    MINIFIED_OUT: 'build.min.js',
    DEST_PUBLIC: 'out/build/public',
    DEST_BUILD: 'out/build',
    TEST_SRC: 'app/js/**/*.js',
    JS_TESTDESTPATH:'out/js-temp',
    JS_TEST_FILES:'out/js-temp/test/**/*.test.js'
};



gulp.task('transform', function(){
    browserify({
        entries: ["app/js/src/todoapp.js"],
        transform: [reactify]
    })
        .bundle()
        .pipe(source(path.MINIFIED_OUT))
        .pipe(gulp.dest(path.DEST_PUBLIC));
});

gulp.task('copy', function () {
    gulp.src(path.HTML)
        .pipe(gulp.dest(path.DEST_BUILD));
});



gulp.task('compile-test-js', function(){
    gulp.src(path.TEST_SRC)
    .pipe(babel({ presets: ['react', 'es2015']}))
    .pipe(gulp.dest(path.JS_TESTDESTPATH));
});

gulp.task('clean', function(){
    gulp.src(path.JS_TEST_FILES)
        .pipe(clean());
});

gulp.task('run-test', function(){
    gulp.src(path.JS_TEST_FILES)
        .pipe(mocha({}));
});

gulp.task('build', ['transform', 'copy']);

gulp.task('test', ['clean', 'compile-test-js', 'run-test']);

gulp.task('default', ['build']);