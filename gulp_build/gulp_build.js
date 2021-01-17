const gulp = require('gulp')
const rename = require('gulp-rename')
const shell = require('gulp-shell')
const clean = require('gulp-clean')

const config = require('./build/config')

const imageName = `${config.name}:${config.version}`
const bootDir = `${config.projectHome}/${config.name}`

gulp.task('build:boot', () => {
    return gulp.src(bootDir)
        .pipe(shell(['mvn clean package -DskipTests=true']))
})

gulp.task('build:pre', () => {
    return gulp.src(`${bootDir}/target/${config.name}-1.0.0-SNAPSHOT.jar`)
        .pipe(rename('app.jar'))
        .pipe(gulp.src('./docker/*'))
        .pipe(gulp.src(`${bootDir}/src/main/resources/application-docker.properties`))
        .pipe(gulp.src(`${bootDir}/src/main/resources/logback-spring.xml`))
        .pipe(gulp.dest(`${bootDir}/docker`))
})

gulp.task('build:docker', () => {
    return gulp.src(`${config.dist}/docker`)
        .pipe(shell([`docker build -t ${imageName} ${config.dist}/docker`,
            `docker save -o ${config.dist}/${config.name}.${config.version}.tar ${imageName}`]))
})

gulp.task('docker:clean', () => {
    return gulp.src(config.dist, {read: false, allowEmpty: true})
        .pipe(clean())
})

const defaultTask = (cb) => {
    gulp.series("docker:build")
    cb && cb()
}

exports.docker = gulp.series('docker:clean', 'build:boot', 'build:pre', 'build:docker')
exports.default = defaultTask
