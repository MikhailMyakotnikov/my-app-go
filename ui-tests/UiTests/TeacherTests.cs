using NUnit.Framework;
using OpenQA.Selenium;
using OpenQA.Selenium.Chrome;
using OpenQA.Selenium.Remote;
using OpenQA.Selenium.Support.UI;

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

        string appUrl = "http://localhost:8081/teachers/create";
        driver.Navigate().GoToUrl(appUrl);

        // IWebElement nameInput = driver.FindElement(By.Name("name"));
        // string name = "Преподаватель";
        // nameInput.SendKeys(name);

        // IWebElement submitButton = driver.FindElement(By.CssSelector(
        //     "[data-testid='create-teacher-btn']"));
        // submitButton.Click();

        // WebDriverWait wait = new WebDriverWait(driver, TimeSpan.FromSeconds(5));
        // wait.Until(driver => driver.Url.Contains("/teachers"));

        

        driver.Quit();
    }
}