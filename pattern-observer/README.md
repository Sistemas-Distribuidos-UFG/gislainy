# Padrão Observer

## Intenção

Definir uma dependência um-para-muitos entre objetos, de maneira que quando um objeto muda de estado todos os seus dependentes são notificados e atualizados automaticamente.

## Também conhecido como

- Dependents
- Publish-Subscribe

## Motivação

O padrão Observer descreve como estabelecer esses relacionamentos. Os objetos-chave nesse padrão são _subject_ (assunto) e _observer_ (observador). Um _subject_ pode ter um número qualquer de observadores dependentes. Todos os observadores são notificados quando o _subject_ sofre uma mudança de estado. Em resposta, cada observador requererá o _subject_ para sincronizar o seu estado com o estado do _subject_.

## Aplicabilidade

Use o padrão Observer em qualquer uma das seguintes situações:

- quando uma abstração tem dois aspectos, um dependente do outro. Encapsulando esses aspectos em objetos separados, permite-se variá-los e reutilizá-los independentemente. 
- quando uma mudança em um objeto exige mudanças em outros, e você não sabe quantos objetos necessitam ser mudados.
- quando um objeto deveria ser capaz de notificar outros objetos nem fazer hipóteses, ou usar informações, sobre quem são esses objetos. Em outras palavras, você não quer que esses objetos sejam fortemente acoplados.

## Estrutura

![Estrutura Padrão Observer](images/estrutura-padrão-observer.png)

## Participantes

- *Subject*
  - conhece os seus observadores. Um número qualquer de objetos Observer pode observar um _subject_
  - fornece uma interface para acrescentar e remover objetos, permitindo associar e desassociar objetos observer.
- *Observer*
  - define uma interface de atualização para objetos que deveriam ser notificados sobre mudanças em um *Subject*.
- *ConcreteSubject*
  - armazena estados de interesse para objetos *ConcreteObserver*.
  - envia uma notificação para os seus observadores quando seu estado muda.
- *ConcreteObserver*
  - mantém uma referência para um objeto *ConcreteSubject*.
  - armazena estados que deveriam permanecer consistentes com os do *Subject*.
  - implementa a interface de atualização de Observer, para manter seu estado consistente.

## Consequências

O padrão Observer permite varias _subjects_ e observadores de forma independes. Você pode reutilizar _subjects_ sem reutilizar seus observadores e vice-versa.
Ele permite acrescentar observadores sem modificar o _subjects_ ou outros observadores.

Benefícios adicionais e deficiências do padrão Observer incluem o seguinte:

1. Acoplamento abstrato entre _Subject_ e _Observer_
2. Suporte para comunicações do tipo broadcast
3. Atualizações inesperadas.

## Referência

Esse trecho foi extraído do livro [Padrões de Projetos: Soluções Reutilizáveis de Software Orientados a Objetos](https://en.wikipedia.org/wiki/Design_Patterns).
