using NUnit.Framework;
using OpenQA.Selenium;
using OpenQA.Selenium.Chrome;
using OpenQA.Selenium.Remote;

namespace UiTests;

public class TeacherTests
{
    [Test]
    public void CreateTeacher()
    {
        IWebDriver driver = new ChromeDriver();
        driver.Navigate().GoToUrl("http://host.docker.internal:8081");

        driver.Quit();
    }
}