package contents

const DomainRepositoryTemplate = `
abstract class {{.DomainRepository}} {
  // Define your abstract methods here
}
`

const DataRepositoryTemplate = `
class {{.DataRepository}} implements {{.DomainRepository}} {
  // Implement your methods here
}
`
