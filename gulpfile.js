var gulp = require('gulp');
var concat = require('gulp-concat');
var uglify = require('gulp-uglify');
var reactify = require('reactify');
var browserify = require('browserify');
var streamify = require('gulp-streamify');

var source = require('vinyl-source-stream');


var path = {
    HTML: 'src/index.html',
    ALL: ['public/scripts/*.js', 'src/index.html'],
    JS: ['public/scripts/*.js'],
    MINIFIED_OUT: 'build.min.js',
    DEST_SRC: 'out/src',
    DEST_PUBLIC: 'out/build/public',
    DEST_BUILD: 'out/build'
};



gulp.task('transform', function(){
    browserify({
        entries: ["src/js/todoapp.js"],
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

gulp.task('build', ['transform', 'copy']);

gulp.task('default', ['build']);