package contents

const UseCaseTemplate = `
class {{.Name}} extends Usecase<{{.ReturnType}}, {{.ParamsType}}> {
  final {{.Repository}} repository;
  {{.Name}}(this.repository);

  @override
  Future<Either<Failure, {{.ReturnType}}>> call({{.ParamsType}} params) async {
    return await repository.{{.Method}}();
  }
}
`
