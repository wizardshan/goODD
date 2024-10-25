# goODD

对象驱动开发，英文全称Object-Driven Development，简称ODD，是一种基于DDD领域驱动设计中的值对象，借鉴Java Optional模式和Google Protobuf3中的Wrapper类型的特性，诞生的一种新型的代码开发模式。

不同于DDD晦涩难懂模棱两可的概念术语，战略战术设计拘谨的学院派风格，ODD只在传统的MVC架构上增加一层属性对象概念，对于开发者更加容易接受和理解，上手难度低，易于推广，代码质量可以显著提高。

#### ODD开发优势：<br>
一、基于Optional模式和Wrapper类型的特性基本杜绝了空指针异常 <br>
二、Set模式解决了数据库增改查零值问题 <br>
三、属性小对象提高了代码可测性 <br>
四、天然的代码内聚性，这也是最重要的一点


- 1、[参数验证如何遵循第一法则DRY？](md/01.md)
- 2、[面向对象编程：万物皆对象](md/02.md)
- 3、[参数的多格式支持](md/03.md)
- 4、[Set模式解决零值问题](md/04.md)
- 5、[Set模式应对查询业务](md/05.md)
- 6、[Set模式应对增删改](md/06.md)
- 7、[gRPC结合ODD开发示例](md/07.md)