var gulp = require('gulp');
var concat = require('gulp-concat');
var uglify = require('gulp-uglify');
var reactify = require('reactify');
var browserify = require('browserify');
var streamify = require('gulp-streamify');
var mocha = require('gulp-mocha');
var babelify = require("babelify");
var babel = require("gulp-babel");

var source = require('vinyl-source-stream');


var path = {
    HTML: 'app/index.html',
    JS: ['app/js/src/*.js'],
    MINIFIED_OUT: 'build.min.js',
    DEST_PUBLIC: 'out/build/public',
    DEST_BUILD: 'out/build',
    TEST_SRC: 'app/js/tests/*.js',
    JS_TESTDESTPATH:'out/build/js-temp'
};



gulp.task('transform', function(){
    browserify({
        entries: ["app/js/src/todoapp.js"],
        transform: [reactify]
    })
        .bundle()
        .pipe(source(path.MINIFIED_OUT))
        //.pipe(streamify(uglify(path.MINIFIED_OUT)))
        .pipe(gulp.dest(path.DEST_PUBLIC));
});

gulp.task('copy', function () {
    gulp.src(path.HTML)
        .pipe(gulp.dest(path.DEST_BUILD));
});

gulp.task('test-compile', function(){
    gulp.src(path.TEST_SRC)
        .pipe(babel())
        .pipe(gulp.dest(path.JS_TESTDESTPATH))
        .pipe(mocha({}));

});

gulp.task('build', ['transform', 'copy']);

gulp.task('default', ['build']);