var gulp = require('gulp');
var concat = require('gulp-concat');
var uglify = require('gulp-uglify');
var react = require('gulp-react');
var htmlreplace = require('gulp-html-replace');


var path = {
    HTML: 'src/index.html',
    ALL: ['public/scripts/*.js', 'src/index.html'],
    JS: ['public/scripts/*.js'],
    MINIFIED_OUT: 'build.min.js',
    DEST_SRC: 'out/src',
    DEST_BUILD: 'out/build',
    DEST: 'out'
};



gulp.task('transform', function(){
    gulp.src(path.JS)
        .pipe(react())
        .pipe(gulp.dest(path.DEST_SRC));
});

gulp.task('copy', function(){
    gulp.src(path.HTML)
        .pipe(gulp.dest(path.DEST));
});

gulp.task('build', ['transform', 'copy']);

gulp.task('default', ['build']);