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

        string seleniumServerUri = "http://localhost:4444";
        IWebDriver driver = new RemoteWebDriver(
            new Uri(seleniumServerUri), 
            options
        );

        string appUrl = "http://host.docker.internal:8081/teachers/create";
        driver.Navigate().GoToUrl(appUrl);

        IWebElement nameInput = driver.FindElement(By.Name("name"));
        string name = "Преподаватель";
        nameInput.SendKeys(name);

        IWebElement submitButton = driver.FindElement(By.CssSelector(
            "[data-testid='create-teacher-btn']"));
        submitButton.Click();

        

        driver.Quit();
    }
}