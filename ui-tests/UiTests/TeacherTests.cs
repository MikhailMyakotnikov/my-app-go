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
        ChromeOptions options = new ChromeOptions();

        IWebDriver driver = new RemoteWebDriver(
            new Uri("http://localhost:4444"), 
            options
        );

        driver.Navigate().GoToUrl("http://localhost:8081");

        driver.Quit();
    }
}